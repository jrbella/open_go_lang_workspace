package main

import(
  f "fmt"
)

func main(){
  var publisher, writer, artist, title string
  var year, pageNumber int
  var grade float32
  
  //Variable Assignments
  title = "Mr. GotoSLeep"
  writer = "Tracy Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  
  f.Println("published by",publisher, "written by",  writer, "artist", artist, "title",title, "year published", year, "total pages",pageNumber, "overall grade", grade)

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0

    f.Println("published by",publisher, "written by",  writer, "artist", artist, "title",title, "year published", year, "total pages",pageNumber, "overall grade", grade)
}
