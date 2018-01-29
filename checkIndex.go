package main

import(

)

/*
* reads hashes from channel checks if, they are in the index
* if not in the index writes them to a new channel
* @param recs, newRecs, dupes : *int
* @return null
*/
func checkIndex( recs, newRecs, dupes *int, index *ind, hashedLineChan, newHashChan *chan string) {

	for hashedLine := range *hashedLineChan {
        if index.contains( hashedLine ){
            *recs++
            *dupes++
        }else{
            *recs++
            *newRecs++
            index.add(hashedLine)
            *newHashChan <- hashedLine
        }
		
	}
	defer close( *newHashChan )
	
}