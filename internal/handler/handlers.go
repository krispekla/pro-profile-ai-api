package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
	"github.com/krispekla/pro-profile-ai-api/internal/repository"
	"github.com/krispekla/pro-profile-ai-api/internal/services"
	"github.com/krispekla/pro-profile-ai-api/types"
)

/*
Package handler provides HTTP request handlers for various endpoints.

For now, all the handlers are implemented in a single file until the file becomes too large.
*/

type Handler struct {
	Db            *sql.DB
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	R2Config      *aws.Config
	CharacterRepo *repository.CharacterRepositoryImpl
	PackageRepo   *repository.PackageRepositoryImpl
	OrderRepo     *repository.OrderRepositoryImpl
	UserRepo      *repository.UserRepositoryImpl
	StripeSvc     *services.StripeServiceImpl
}

func NewHandler(db *sql.DB, errorLog *log.Logger, infoLog *log.Logger, r2Config *aws.Config) *Handler {
	characterRepo := repository.NewCharacterRepositoryImpl(db)
	packageRepo := repository.NewPackageRepositoryImpl(db)
	orderRepo := repository.NewOrderRepositoryImpl(db)
	userRepo := repository.NewUserRepositoryImpl(db)
	stripeSvc := services.NewStripeServiceImpl(orderRepo)
	return &Handler{
		Db:            db,
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		R2Config:      r2Config,
		CharacterRepo: characterRepo,
		PackageRepo:   packageRepo,
		OrderRepo:     orderRepo,
		UserRepo:      userRepo,
		StripeSvc:     stripeSvc,
	}
}

// TODO: For all errors return appropriate client status and message

func (h *Handler) UserDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ErrorLog.Print("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

type UserRow struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserWebhookPayload struct {
	Type   string  `json:"type"`
	Table  string  `json:"table"`
	Record UserRow `json:"record"`
}

func (h *Handler) UserRegistrationWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const MaxBodyBytes = int64(65536)
		r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			h.ErrorLog.Print("Error reading request body")
			return
		}
		var data UserWebhookPayload
		err = json.Unmarshal(payload, &data)
		if err != nil {
			h.ErrorLog.Print("Error unmarshaling request body")
			return
		}
		// Create Stripe customer for user
		usrId, err := uuid.Parse(data.Record.ID)
		if err != nil {
			h.ErrorLog.Print("Error parsing uuid")
			return
		}
		fullName := ""
		if data.Record.FirstName != "" {
			fullName = data.Record.FirstName
		}
		if data.Record.LastName != "" {
			if fullName != "" {
				fullName += " "
			}
			fullName += data.Record.LastName
		}
		customerInput := &services.CreateCustomerInput{
			UserId:   data.Record.ID,
			Email:    data.Record.Email,
			FullName: fullName,
		}
		cstmr, err := h.StripeSvc.CreateCustomer(customerInput)
		if err != nil {
			h.ErrorLog.Print("Error creating customer")
			return
		}
		err = h.UserRepo.UpdateCustomerDetails(&repository.UserCustomerInput{Id: usrId, StripeCustomerID: cstmr.ID})
		if err != nil {
			h.ErrorLog.Print("Error updating user with customer details")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) GetCharacters() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.Context().Value(types.UserContextKey).(*types.JwtUser)
		usrId, err := uuid.Parse(usr.Id)
		if err != nil {
			h.ErrorLog.Print("Error parsing uuid")
		}
		// Check to see if
		result, err := h.CharacterRepo.Get(usrId)
		if err != nil {
			h.ErrorLog.Print("Error retrieving characters")
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			h.ErrorLog.Print("Error marshaling characters")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	}
}

func (h *Handler) GetPackageListing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.PackageRepo.GetListing()
		if err != nil {
			h.ErrorLog.Print("Error retrieving packages")
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			h.ErrorLog.Print("Error marshaling packages")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	}
}

func (h *Handler) GetGeneratedPackages() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.Context().Value(types.UserContextKey).(*types.JwtUser)
		usrId, err := uuid.Parse(usr.Id)
		if err != nil {
			h.ErrorLog.Print("Error parsing uuid")
		}
		result, err := h.PackageRepo.GetGeneratedPackages(usrId)
		if err != nil {
			h.ErrorLog.Print("Error retrieving packages")
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			h.ErrorLog.Print("Error marshaling packages")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	}
}

func (h *Handler) GetAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.Context().Value(types.UserContextKey).(*types.JwtUser)
		usrId, err := uuid.Parse(usr.Id)
		if err != nil {
			h.ErrorLog.Print("Error parsing uuid")
		}
		result, err := h.OrderRepo.GetAllOrders(usrId)
		if err != nil {
			h.ErrorLog.Print("Error retrieving packages")
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			h.ErrorLog.Print("Error marshaling packages")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	}
}

func (h *Handler) GetPackageDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name": "Package 1", "price": 30}, images": ["image1.jpg", "image2.jpg"], "description": "This is a package description"`))
	}
}

func (h *Handler) BuyPackage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Package bought"}`))
	}
}

func (h *Handler) CreateCharacter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.Context().Value(types.UserContextKey).(*types.JwtUser)
		usrId, err := uuid.Parse(usr.Id)
		if err != nil {
			h.ErrorLog.Print("Error parsing uuid")
		}
		chrNew, err := h.CharacterRepo.CreateCharacter(usrId)
		if err != nil {
			h.ErrorLog.Print("Error creating character")
			return
		}
		chrJson, err := json.Marshal(chrNew)
		if err != nil {
			h.ErrorLog.Print("Error marshaling character")
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(chrJson)
	}
}

func (h *Handler) UpdateCharacter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body repository.UpdateCharacterInput
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			h.ErrorLog.Print("Error decoding character body")
			return
		}
		chrNew, err := h.CharacterRepo.UpdateCharacter(&body)
		if err != nil {
			h.ErrorLog.Print("Error updating character")
			return
		}
		chrJson, err := json.Marshal(chrNew)
		if err != nil {
			h.ErrorLog.Print("Error marshaling character")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(chrJson)
	}
}

type CreateCheckoutSessionBody struct {
	StripeProductIds *[]string `json:"productIds"`
}

func (h *Handler) CreateCheckoutSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body CreateCheckoutSessionBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			h.ErrorLog.Print("Error decoding checkout body")
			return
		}
		if body.StripeProductIds == nil || len(*body.StripeProductIds) == 0 {
			h.ErrorLog.Print("ProductIds are required")
			return
		}
		usrCtxId := r.Context().Value(types.UserContextKey).(*types.JwtUser)
		usrId, err := uuid.Parse(usrCtxId.Id)
		if err != nil {
			h.ErrorLog.Print("Error parsing uuid")
			return
		}
		usr, err := h.UserRepo.Get(usrId)
		if err != nil {
			h.ErrorLog.Print("Error retrieving user")
			return
		}
		var customerId string
		if usr.StripeCustomerID == nil || *usr.StripeCustomerID == "" {
			fullName := ""
			if usr.FirstName != nil {
				fullName = *usr.FirstName
			}
			if usr.LastName != nil {
				if fullName != "" {
					fullName += " "
				}
				fullName += *usr.LastName
			}
			if usrCtxId.Email == "" {
				h.ErrorLog.Print("Error retrieving user email")
				return
			}
			// Create customer
			customerInput := &services.CreateCustomerInput{
				UserId:   usrCtxId.Id,
				Email:    *usr.Email,
				FullName: fullName,
			}
			cstmr, err := h.StripeSvc.CreateCustomer(customerInput)
			if err != nil {
				h.ErrorLog.Print("Error creating customer")
				return
			}
			customerId = cstmr.ID
			err = h.UserRepo.UpdateCustomerDetails(&repository.UserCustomerInput{Id: usrId, StripeCustomerID: customerId})
			if err != nil {
				h.ErrorLog.Print("Error updating user with customer details")
				return
			}
		} else {
			customerId = *usr.StripeCustomerID
		}
		// Get price id from request
		checkoutParam := &services.CreateCheckoutSessionInput{ProductIds: body.StripeProductIds, CustomerId: customerId}
		s, err := h.StripeSvc.CreateCheckoutSession(checkoutParam)

		if err != nil {
			h.ErrorLog.Print("Error creating checkout session")
			return
		}

		pprices, err := h.PackageRepo.GetPackagePrice(body.StripeProductIds)
		if err != nil {
			h.ErrorLog.Print("Error retrieving package prices")
			return
		}

		oi := &repository.CreateOrderInput{
			CheckoutId:    &s.ID,
			Amount:        &s.AmountTotal,
			Currency:      string(s.Currency),
			UserId:        &usrId,
			PackagePrices: pprices,
		}

		oCreated, err := h.OrderRepo.CreateOrder(oi)
		if err != nil {
			h.ErrorLog.Print("Error creating order")
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Checkout session created", "clientSecret": "` + s.ClientSecret + `" , "orderId": "` + string(oCreated.ID) + `"}`))
	}
}

func (h *Handler) RetrieveCheckoutSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.URL.Query().Get("session_id")
		s, err := h.StripeSvc.GetStripeSession(sessionId)
		if err != nil {
			h.ErrorLog.Print("Error retrieving checkout session")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Checkout session retrieved", "status": "` + string(s.Status) + `", "customerEmail": "` + string(s.CustomerDetails.Email) + `"}`))
	}
}
func (h *Handler) StripeWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const MaxBodyBytes = int64(65536)
		r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		stripeSign := r.Header.Get("Stripe-Signature")
		status, err := h.StripeSvc.ProcceesStripeWebhook(payload, stripeSign)
		if err != nil {
			h.ErrorLog.Print(err)
		}
		w.WriteHeader(status)
	}
}

// TODO: Separate R2 config and storage config in upper level

// func (h *Handler) GetAllBuckets() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		client := s3.NewFromConfig(*h.R2Config)

// 		listObjectsOutput, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
// 			Bucket: &h.StorageConfig.R2BucketName,
// 		})
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		for _, object := range listObjectsOutput.Contents {
// 			obj, _ := json.MarshalIndent(object, "", "\t")
// 			fmt.Println(string(obj))
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message": "Buckets retrieved"}`))

// 		//  {
// 		//  	"ChecksumAlgorithm": null,
// 		//  	"ETag": "\"eb2b891dc67b81755d2b726d9110af16\"",
// 		//  	"Key": "ferriswasm.png",
// 		//  	"LastModified": "2022-05-18T17:20:21.67Z",
// 		//  	"Owner": null,
// 		//  	"Size": 87671,
// 		//  	"StorageClass": "STANDARD"
// 		//  }

// 		// listBucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
// 		// if err != nil {
// 		// 	log.Fatal(err)
// 		// }

// 		// for _, object := range listBucketsOutput.Buckets {
// 		// 	obj, _ := json.MarshalIndent(object, "", "\t")
// 		// 	fmt.Println(string(obj))
// 		// }

// 		// {
// 		// 		"CreationDate": "2022-05-18T17:19:59.645Z",
// 		// 		"Name": "sdk-example"
// 		// }
// 	}
// }

// func (h *Handler) GetPresignedImgUrl() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		client := s3.NewFromConfig(*h.R2Config)
// 		presignClient := s3.NewPresignClient(client)

// 		key := "kris555.jpg"

// 		// presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
// 		// 	Bucket: &h.R2BucketName,
// 		// 	Key:    &key,
// 		// })

// 		presignResult, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
// 			Bucket: &h.StorageConfig.R2BucketName,
// 			Key:    &key,
// 		})

// 		if err != nil {
// 			panic("Couldn't get presigned URL for GetObject")
// 		}

// 		fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)
// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte(`{"message": "Presigned URL retrieved", "url": "` + presignResult.URL + `"}`))
// 	}
// }
