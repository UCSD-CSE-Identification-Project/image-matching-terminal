package main

import (
  "database/sql"
  "fmt"
  "bufio"
  "os"
  "strings"
  "log"
)

type  uploadData struct {
  Link  string  `json:"link"`
}

type imageObject struct {
  key             string
  hash_val        string
  ocr_text        string
  strict_match    string
  tier_one        string
  tier_two        string

}

const (
  KEY_CONST = "KEY"
  HASH_VAL_CONST = "HASH_VAL"
  OCR_TEXT_CONST = "OCR_TEXT"
  STRICT_MATCH_CONST = "STRICT_MATCH_NAMES"
  POTENTIAL_MATCH_CONST = "POTENTIAL_MATCH_NAMES"
  TIER_ONE_CONST = "Tier 1:"
  TIER_TWO_CONST = "Tier 2:"
)
func (u *uploadData) uploadToDatabase(db *sql.DB){

  imageObjects:=parseFile(u.Link)
  populatedDB(db, imageObjects)
}

func parseFile(link string) []imageObject {

  file, err := os.Open(link) // open the file
  if err != nil {
    fmt.Println(err)
    return nil
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  //scanner.Split(bufio.ScanRunes)
  scanner.Split(bufio.ScanLines)

  // This is our buffer now
  var lines []string
  var currentLine string
  //var currentLine string
  //var lines []rune
  var curImageObj imageObject
  var imageObjects []imageObject

  for scanner.Scan() {
    currentLine = scanner.Text()
    switch str := doesContainKeyword(currentLine); str {
    case KEY_CONST:
      lines = append(lines,currentLine[5:])
      curImageObj.key = currentLine[5:]
    case HASH_VAL_CONST:
      scanner.Scan()
      hashval:= scanner.Text()[8:]
      lines = append(lines, hashval)
      curImageObj.hash_val = hashval
    case OCR_TEXT_CONST:
      scanner.Scan()
      ocrTextVal := scanner.Text()
      lines = append(lines,ocrTextVal )
      curImageObj.ocr_text = ocrTextVal
    case STRICT_MATCH_CONST:
      scanner.Scan()
      strictMatch := scanner.Text()
      lines = append(lines, strictMatch)
      curImageObj.strict_match = strictMatch

    case POTENTIAL_MATCH_CONST:
      continue
    case TIER_ONE_CONST:
      scanner.Scan()
      tierOne := scanner.Text()
      lines = append(lines, tierOne)
      curImageObj.tier_one = tierOne
    case TIER_TWO_CONST:
      scanner.Scan()
      tierTwo := scanner.Text()
      lines = append(lines, tierTwo)
      curImageObj.tier_two = tierTwo

      //printObj(curImageObj)
      imageObjects = append(imageObjects, curImageObj)
      curImageObj = imageObject{}

    default:
      log.Println("==============")
      //log.Println(currentLine)
      //log.Println("Rohan")
    }
    //lines = append(lines, currentLine)
    //currentLine = scanner.Text()
  }

  //log.Println(imageObjects)
  return imageObjects

}

func doesContainKeyword(scanned string) string{

  if strings.Contains(scanned,KEY_CONST){
    return KEY_CONST
  }else if strings.Contains(scanned,HASH_VAL_CONST) {
    return HASH_VAL_CONST
  }else if strings.Contains(scanned,OCR_TEXT_CONST) {
    return OCR_TEXT_CONST
  }else if strings.Contains(scanned,STRICT_MATCH_CONST) {
    return STRICT_MATCH_CONST
  }else if strings.Contains(scanned,POTENTIAL_MATCH_CONST) {
    return POTENTIAL_MATCH_CONST
  }else if strings.Contains(scanned, TIER_ONE_CONST) {
    return TIER_ONE_CONST
  }else if strings.Contains(scanned, TIER_TWO_CONST) {
    return TIER_TWO_CONST
  }else{
    return "none"
  }
}

func printObj( obj imageObject){
  log.Println(obj.key, obj.hash_val, obj.strict_match, obj.tier_one, obj.tier_two)
}
func printObjs( objs []imageObject ){
  for _, element := range objs {
    // element is the element from someSlice for where we are
    printObj(element)
  }
}

func populatedDB( db *sql.DB, imgO []imageObject ) {
  for _, element := range imgO {
    // element is the element from someSlice for where we are
    //printObj(element)
    //if index > 10 {
    //  continue
    //}
    db.Exec("INSERT INTO image_info(key,hash_val,ocr_text) VALUES($1,$2,$3)", element.key,element.hash_val,element.ocr_text)

    for _, t1Elem := range strings.Split(element.tier_one, " ") {
      if strings.Compare(t1Elem,"") == 0 || strings.Compare(t1Elem,"\t") == 0 {
        continue
      }
      t1Elem = strings.Replace(t1Elem,"\t","",-1)
      //log.Println(t1Elem)
      db.Exec("INSERT INTO image_tier_one(key,match) VALUES($1,$2)", element.key,t1Elem)

    }
    for _, t2Elem := range strings.Split(element.tier_two, " ") {
      if strings.Compare(t2Elem,"") == 0 || strings.Compare(t2Elem,"\t") == 0 {
        continue
      }
      t2Elem = strings.Replace(t2Elem,"\t","",-1)
      //log.Println(t2Elem)
      db.Exec("INSERT INTO image_tier_two(key,match) VALUES($1,$2)", element.key,t2Elem)
    }
    for _, strictElem := range strings.Split(element.strict_match, " ") {
      if strings.Compare(strictElem,"") == 0 || strings.Compare(strictElem,"\t") == 0 {
        continue
      }
      strictElem = strings.Replace(strictElem,"\t","",-1)
      //log.Println(strictElem)
      db.Exec("INSERT INTO strict_match(key,match) VALUES($1,$2)", element.key,strictElem)
    }
  }
  //_, err := db.Exec("INSERT INTO logincredentials(username,password) VALUES($1,$2)", l.UserName, l.Password)

}




