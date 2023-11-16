package main

import "mistakes/9_embedding/store"

func main() {
	st := store.NewStore()
	st.Unlock()

	st2 := store.NewStore2()
	st2.Get()
}
