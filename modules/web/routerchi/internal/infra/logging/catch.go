package logging

func Catch(err error) {
	if err != nil {
		GetLogger().Panic(err)
		panic(err)
	}
}
