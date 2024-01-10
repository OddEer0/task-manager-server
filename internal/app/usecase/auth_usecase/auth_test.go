package authUsecase

//func TestAuthUseCaseRegistration(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	// Создаем мок AuthUseCase
//	mockAuthUseCase := mocks.NewMockAuthUseCase(ctrl)
//
//	// Задаем ожидаемое поведение для метода Registration
//	user := &models.User{ID: 1, Username: "testuser"}
//	mockAuthUseCase.EXPECT().Registration(gomock.Any(), gomock.Any()).Return(user, nil)
//
//	// Используем мок в тесте
//	ctx := context.TODO()
//	registrationData := dto.RegistrationInputDto{Username: "testuser", Password: "testpassword"}
//	result, err := mockAuthUseCase.Registration(ctx, registrationData)
//
//	// Проверяем, что ожидаемое поведение было вызвано
//	// Этот вызов будет проходить успешно, если метод Registration был вызван ожидаемым образом
//
//	// Проверяем результат
//	assert.NoError(t, err)
//	assert.NotNil(t, result)
//	assert.Equal(t, user, result)
//}
