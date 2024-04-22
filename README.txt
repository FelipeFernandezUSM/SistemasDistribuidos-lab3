Felipe Ignacio Fernandez Aguilar ROL:201856602-0
Christian Renato Ossa Salgado    ROL:201873565-5

Uso del codigo:

Para utilizar el codigo desarrollado se deben seguir los siguientes pasos

1.- abrir 5 terminales y cambiar el directorio de una a la carpeta server_central y las otras a soldier
2.- una vez en server_central se debe correr el comando por terminal docker build -t server_central ., si lo buscan correr denuevo este paso se puede saltar (cuando queda bien creada la imagen)
3.- una vez hecho esto podemos correr el docker con el comando por terminal docker run -p 8080:8080 server_central
4.- ahora las terminales soldier se correra el go run main.go en todas lo mas cercano posible (al mismo tiempo seria lo mejor)

estos son los pasos necesarios para la ejecucion del codigo