package internal

func Sphere() Shape {
	return Shape{
		Transform: IdentityMatrix(),
		Material:  NewMaterial(),
	}
}
