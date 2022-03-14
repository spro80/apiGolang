package orderUseCase

import (
	"fmt"

	"github.com/spro80/apiGolang/app/application"
)

type OrderUseCaseInterface interface {
	application.UseCase
	Process() (ResponseProcessOrderUseCase, error)
}

type orderUseCase struct {
}

type ResponseProcessOrderUseCase struct {
	Status            bool   `json:"status"`
	StatusCode        int    `json:"statusCode"`
	StatusDescription string `json:"statusDescription"`
	StatusError       string `json:"statusError"`
}

func NewOrderUseCase() *orderUseCase {
	return &orderUseCase{}
}

//func OrderUseCase() (bool, error) {
func (o *orderUseCase) Process() (ResponseProcessOrderUseCase, error) {
	fileName := "orderUseCase"
	fmt.Printf("\n[%s] OrderUseCase Init in orderUseCase", fileName)

	var id int = 100
	var name string = "David"

	status, err := SaveInDatabase(id, name)
	if err != nil {
		fmt.Printf("\n[%s] OrderUseCase method save was called with error", fileName)
		res := ResponseProcessOrderUseCase{
			Status:            false,
			StatusCode:        0,
			StatusDescription: "Error in call method",
			StatusError:       "",
		}
		return res, nil
	}
	fmt.Printf("\n[%s] OrderUseCase method save was called succesfully", fileName)
	fmt.Println(status)

	res := ResponseProcessOrderUseCase{
		Status:            true,
		StatusCode:        1,
		StatusDescription: "Hello World Use case",
		StatusError:       "",
	}

	return res, nil
}

func SaveInDatabase(id int, name string) (bool, error) {
	fmt.Println("[orderUseCase] saveInDatabase")
	fmt.Println("[orderUseCase] saveInDatabase id:")
	fmt.Println(id)
	fmt.Println("[orderUseCase] saveInDatabase name:")
	fmt.Println(name)
	return true, nil
}

/*
func NewOrderUseCase(repository repository.OrdersRepository) *orderUseCase {
	return &orderUseCase{
		ordersRepository: repository,
	}
}
*/

/*
//func (s *orderUseCase) Process(order *entity.Order) (*entity.Order, error) {
func (s *orderUseCase) Process(order *entity.Order) (*entity.Order, error) {
	fmt.Println("[orderUseCase] Process - step 1")
	order.State = status.INIT
	savedOrder, err := s.ordersRepository.Create(order)
	if err != nil {
		fmt.Println("[orderUseCase] Process - step 2")
		fmt.Println("[orderUseCase] Error != nil")
		//return nil, 500
		//log.WithError(err).Error("Error trying to save order")
		//return nil, custom_errors.NewWithError(err, custom_errors.DataBaseError)
	}

	//log.WithFields(log.Field("order", savedOrder.Id)).Info("Order saved successfully")
	//metrics.IncrementTestMetrics("SAVE")
	return savedOrder, nil

*/
/*
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex( "6205707eae686ee4857130bc")

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	fmt.Println("[orderUseCase] Process - step 3 - success")
	return savedOrder, nil
*/
//}
