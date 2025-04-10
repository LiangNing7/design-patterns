package factorymethod

func ExampleFactoryMethod() {
	FactoryMethod(&FileLoggerFactory{})
	FactoryMethod(&ConsoleLoggerFactory{})

	// Output:
	// Log to file: This is a test log.
	// Log to console: This is a test log.
}

func FactoryMethod(factory LoggerFactory) {
	factory.CreateLogger().Log("This is a test log.")
}
