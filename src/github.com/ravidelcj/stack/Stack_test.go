package stack

import "testing"



func Test(t *testing.T)  {
    elem := element{title : "title", date : "date", url : "url", remoteUrl : "remoteUrl"}
    var st StackNode

    //checking initial top position
    _, err := st.top()

    if err == true {
      t.Error("Error in first top")
    }

    //initial pop checking
    popElement := st.pop()
    if popElement == true {
        t.Error("error in first pop")
    }

    //testing initial empty
    s := st.isEmpty()
    if s == false {
      t.Error("empty")
    }

    //initial size checking
    size1 := st.size()
    if size1 != 0 {
      t.Error("size not equal 0")
    }

    //pushing one element
    v := st.push(elem)
    if v == false {
      t.Error("error in push")
    }

    //testing size
    size2 := st.size()
    if size2 != 1 {
      t.Error("size not equal 1")
    }

    //pushing second element
    u := st.push(elem)
    if u == false {
      t.Error("error in push")
    }

    //getting top element
    topElement, err := st.top()

    if err == false {
      t.Error("Error in first top")
    }
    if topElement != elem {
      t.Error("Error in retreiving")
    }

    //size testing
    size3 := st.size()
    if size3 != 2 {
      t.Error("size not equal 1")
    }

    //testing empty condition
    q := st.isEmpty()
    if q == true {
      t.Error("empty")
    }
}
