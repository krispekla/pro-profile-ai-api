package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	authTable "github.com/krispekla/pro-profile-ai-api/.gen/postgres/auth/table"
	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/model"
	. "github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/table"
	"github.com/krispekla/pro-profile-ai-api/config"
	"github.com/krispekla/pro-profile-ai-api/types"
	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

// func login(app *config.Application) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Get the username and password from the request body
// 		email := r.PostFormValue("email")
// 		password := r.PostFormValue("password")
// 		// Check if the email and password are valid

// 		if email == "" || password == "" {
// 			app.ClientError(w, http.StatusBadRequest)
// 			return
// 		}

// 		msg := fmt.Sprintf("User %s is logged in", email)

// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(msg))
// 	}
// }

// func register(app *config.Application) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		firstName := r.PostFormValue("first_name")
// 		lastName := r.PostFormValue("last_name")
// 		email := r.PostFormValue("email")
// 		password := r.PostFormValue("password")

// 		if firstName == "" || lastName == "" || email == "" || password == "" {
// 			app.ClientError(w, http.StatusBadRequest)
// 			return
// 		}

// 		w.WriteHeader(http.StatusCreated)
// 		uCreatedMsg := fmt.Sprintf("User created: %s %s", firstName, lastName)
// 		w.Write([]byte(uCreatedMsg))
// 	}
// }

func userDetails(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.ErrorLog.Print("Unauthorized")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func getCharacters(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := r.Context().Value(types.UserContextKey).(*types.JwtUser)
		stmt := SELECT(
			Character.AllColumns,
			authTable.Users.Email.AS("email"),
		).FROM(Character.LEFT_JOIN(
			authTable.Users, Character.UserID.EQ(UUID(uuid.MustParse(usr.Id)))))

		var result []struct {
			model.Character
			email string `db:"email"`
		}

		err := stmt.Query(app.Db, &result)
		if err != nil {
			app.ErrorLog.Print("Error retrieving characters")
		}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			app.ErrorLog.Print("Error marshaling characters")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	}
}

func getPackages(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"packages": [{"name": "Package 1", "price": 30}, {"name": "Package 2", "price": 25}]}`))
	}
}

func getPackageDetails(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name": "Package 1", "price": 30}, images": ["image1.jpg", "image2.jpg"], "description": "This is a package description"`))
	}
}

func buyPackage(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Package bought"}`))
	}
}

func createCharacter(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Character created"}`))
	}
}

func createCheckoutSession(app *config.Application) http.HandlerFunc {
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
			app.ErrorLog.Print("Error creating checkout session")
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Checkout session created", "clientSecret": "` + s.ClientSecret + `"}`))
	}
}

func retrieveCheckoutSession(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, _ := session.Get(r.URL.Query().Get("session_id"), nil)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Checkout session retrieved", "status": "` + string(s.Status) + `", "customerEmail": "` + string(s.CustomerDetails.Email) + `"}`))
	}
}

func getAllBuckets(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		client := s3.NewFromConfig(*app.R2Config)

		listObjectsOutput, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: &app.StorageConfig.R2BucketName,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, object := range listObjectsOutput.Contents {
			obj, _ := json.MarshalIndent(object, "", "\t")
			fmt.Println(string(obj))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Buckets retrieved"}`))

		//  {
		//  	"ChecksumAlgorithm": null,
		//  	"ETag": "\"eb2b891dc67b81755d2b726d9110af16\"",
		//  	"Key": "ferriswasm.png",
		//  	"LastModified": "2022-05-18T17:20:21.67Z",
		//  	"Owner": null,
		//  	"Size": 87671,
		//  	"StorageClass": "STANDARD"
		//  }

		// listBucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// for _, object := range listBucketsOutput.Buckets {
		// 	obj, _ := json.MarshalIndent(object, "", "\t")
		// 	fmt.Println(string(obj))
		// }

		// {
		// 		"CreationDate": "2022-05-18T17:19:59.645Z",
		// 		"Name": "sdk-example"
		// }
	}
}

func getPresignedImgUrl(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client := s3.NewFromConfig(*app.R2Config)
		presignClient := s3.NewPresignClient(client)

		key := "kris555.jpg"

		// presignResult, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		// 	Bucket: &app.R2BucketName,
		// 	Key:    &key,
		// })

		presignResult, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: &app.StorageConfig.R2BucketName,
			Key:    &key,
		})

		if err != nil {
			panic("Couldn't get presigned URL for GetObject")
		}

		fmt.Printf("Presigned URL For object: %s\n", presignResult.URL)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Presigned URL retrieved", "url": "` + presignResult.URL + `"}`))
	}
}
