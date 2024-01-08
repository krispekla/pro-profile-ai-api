package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
	"github.com/krispekla/pro-profile-ai-api/internal/repository"
	"github.com/krispekla/pro-profile-ai-api/types"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
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
}

func NewHandler(db *sql.DB, errorLog *log.Logger, infoLog *log.Logger, r2Config *aws.Config) *Handler {
	characterRepo := repository.NewCharacterRepositoryImpl(db)
	packageRepo := repository.NewPackageRepositoryImpl(db)
	orderRepo := repository.NewOrderRepositoryImpl(db)
	return &Handler{
		Db:            db,
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		R2Config:      r2Config,
		CharacterRepo: characterRepo,
		PackageRepo:   packageRepo,
		OrderRepo:     orderRepo,
	}
}

func (h *Handler) UserDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ErrorLog.Print("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Character created"}`))
	}
}

func (h *Handler) CreateCheckoutSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := "http://localhost:5173"
		params := &stripe.CheckoutSessionParams{
			UIMode:    stripe.String("embedded"),
			ReturnURL: stripe.String(domain + "/package/buy/return?session_id={CHECKOUT_SESSION_ID}"),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				&stripe.CheckoutSessionLineItemParams{
					// TODO: Product id from stripe dashboard, adjust for production
					Price:    stripe.String("price_1OITFBFSEa3MNRY9u9T0IOy6"),
					Quantity: stripe.Int64(1),
				},
			},
			Mode:         stripe.String(string(stripe.CheckoutSessionModePayment)),
			AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{Enabled: stripe.Bool(true)},
		}

		s, err := session.New(params)

		if err != nil {
			h.ErrorLog.Print("Error creating checkout session")
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Checkout session created", "clientSecret": "` + s.ClientSecret + `"}`))
	}
}

func (h *Handler) RetrieveCheckoutSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := session.Get(r.URL.Query().Get("session_id"), nil)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Checkout session retrieved", "status": "` + string(s.Status) + `", "customerEmail": "` + string(s.CustomerDetails.Email) + `"}`))
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
