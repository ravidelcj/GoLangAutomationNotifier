package stack

import "testing"



func Test(t *testing.T)  {
    elem := Element{Title : "title", Date : "date", Url : "url", RemoteUrl : "remoteUrl"}
    st := NewStack()

    //checking initial top position
    _, err := st.Top()

    if err == true {
      t.Error("Error in first top")
    }

    //initial pop checking
    popElement := st.Pop()
    if popElement == true {
        t.Error("error in first pop")
    }

    //testing initial empty
    s := st.IsEmpty()
    if s == false {
      t.Error("empty")
    }

    //initial size checking
    size1 := st.Size()
    if size1 != 0 {
      t.Error("size not equal 0")
    }

    //pushing one element
    v := st.Push(elem)
    if v == false {
      t.Error("error in push")
    }

    //testing size
    size2 := st.Size()
    if size2 != 1 {
      t.Error("size not equal 1")
    }

    //pushing second element
    u := st.Push(elem)
    if u == false {
      t.Error("error in push")
    }

    //getting top element
    topElement, err := st.Top()

    if err == false {
      t.Error("Error in first top")
    }
    if topElement != elem {
      t.Error("Error in retreiving")
    }

    //size testing
    size3 := st.Size()
    if size3 != 2 {
      t.Error("size not equal 1")
    }

    //testing empty condition
    q := st.IsEmpty()
    if q == true {
      t.Error("empty")
    }
}
