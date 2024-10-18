package main
import ("fmt")

func main() {
	// way to init map
	m1 := map[string]string{"one":"1", "two":"2"};

	fmt.Println("m1: ", m1)

	var m2 = map[string]string{"xyz":"abc","abc":"xyz"}
	fmt.Println(m2)

	// map key type only those over which equality (==) can be applied, pretty much only primitive types

	// empty map
	// two ways
	// 1. by make() or :=
	// 2, by var 
	// problem with 2nd is that on updation would panic right away

	m3 := make(map[string]string)
	fmt.Println(m3)
	m3["name"] = "sdkjcn"

	fmt.Println(m3)

	m4 := map[string]string{}
	m4["name"] = "sdkjcn"
	fmt.Println(m4)
	var m5 map[string]string
	// m5["name"] = "sdkjcn" // this would panic

	fmt.Println(m5)

	// delete elem from map
	delete(m1, "one")
	fmt.Println(m1)

	// check for existance for a elem in map
	itsme := map[string]string{"name":"abdus","rollno":"1223","city":"rrk"}

	_, ok := itsme["one"]
	val, ok1 := itsme["name"]
	_, ok2 := itsme["city"]
	// _ means do not wanna use it, this would not result in variable naming colliosn like in js or ts
	fmt.Println(ok)
	fmt.Println(val, ok1)
	fmt.Println( ok2)

	// map are ref to hash table
	// if two variable pointing to same map hash table
	// changing map using one var, reflect on other too

	itsmeCopy := itsme

	itsmeCopy["name"] = "abdus samad"
	fmt.Println(itsme)
	fmt.Println(itsmeCopy)

	for i:=0; i < 10; i++ {
		key := fmt.Sprintf("key%v", i)
		itsme[key] = key;
	}

	// iteration	
	fmt.Println("looping with no order")
	for k,v:= range itsme {
		if k =="city" {
			break;
		}
		fmt.Println(k ,v)
	}

	// appr to acces in an order
	keysOrder := []string{"name", "city", "rollno"};
	for _, elem := range keysOrder {
		fmt.Println(elem,  " => ", itsme[elem])
	} 
//   var a = make(map[string]string)
//   var b map[string]string


// 	m := map[string]string{"hello":"csd"};

//   fmt.Println(a == nil)
//   fmt.Println(b == nil)
// 	m["hello1"]= "csdjk";
	
// 	fmt.Println(m)
// 	_, ok1 := m["hello"]
// 	_, ok2 := m["hello"]
// 	fmt.Println(ok1)
// 	fmt.Println(ok2)

}