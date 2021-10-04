► El metodo Parser recibe la entidad y devuelve e informa si hay o no error.

→ Como primer paso se eliminan los espacios de la misma.

→ Guardo el tipo de la entidad. (Pos 0 y 1 de la entidad) 

→ Guardo el tamaño declarado. Si empieza con 0, devuelvo el siguiente y sino, lo devuelvo completo.

→ Guardo los valores apartir de la pos 4 y a su vez, el length desde dicha pos.

→ Identifico el tipo de la entidad y si el mismo concuerda con los valores que fueron insertados posteriormente, advirtiendo si hay un error o no.

→ Parseo el tamaño de la entidad a int ya que antes era un string.

→ En este punto hago condiciones para detectar los errores que se puedan detectar.
