Resolución

► El metodo Parser recibe la entidad y devuelve e informa si hay o no error.

→ Como primer paso se eliminan los espacios de la misma.

→ Guardo el tipo de la entidad. (Pos 0 y 1 de la entidad) 

→ Guardo el tamaño declarado. Si empieza con 0, devuelvo el siguiente y sino, lo devuelvo completo.

→ Guardo los valores apartir de la pos 4 y a su vez, el length desde dicha pos.

→ Identifico el tipo de la entidad y si el mismo concuerda con los valores que fueron insertados posteriormente, advirtiendo si hay un error o no.

→ Parseo el tamaño de la entidad a int ya que antes era un string.

→ En este punto hago condiciones para detectar los errores que se puedan detectar.

------------------------------------------------------------------------------------------------------------------------------------------------

Enunciado

Crear una funcion que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena. Por ejemplo:

TX03ABC
Deberá generar una estructura con los siguientes valores:

{TX 3 ABC}
Donde la estructura deberá tener una definicion del tipo:

type Result struct {
    Type    string 
    Value   string
    Length  int
}
La cadena de caracteres tiene el siguiente formato:

los primeros dos caracteres son el tipo
los segundos dos caracteres son el largo del valor
los siguientes N caracteres son el valor, donde N es el valor del bullet anterior.
Entonces NN040001 equivale a:

Type=NN
Length=04
Value=0001
