package main

func main() {
	runner(map[string]func(){
		"env": demo_env,
	})
}
