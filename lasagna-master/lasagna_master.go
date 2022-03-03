package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, x int) int {
	if x == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * x

	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int , float64){
	noodles := 0.0
	lasagna := 0.0
	for _, v := range layers {
		if(v == "noodles"){
			noodles +=50
		}else if v == "sauce" {
			lasagna += 0.2
		}
		
	}
	return int(noodles),lasagna
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend []string,my []string,)  []string{
	
	return append(my, friend[len(friend)-1])
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(r []float64,x int)  []float64{
	
	y := make([]float64, len(r))

	for i, v := range r{
		//https://stackoverflow.com/questions/32815400/how-to-perform-division-in-go
		y[i] = v * (float64(x)/2)
	}

	return y
}