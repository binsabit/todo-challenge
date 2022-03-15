console.log("JS")
const URL = "http://localhost:3000/todo"
var inputForm = document.querySelector("button")

showList()

inputForm.addEventListener("click",async (e)=>{
    e.preventDefault()
    inp = document.querySelector("#inp")
    let data = {
        itemName : inp.value
    }
    console.log(data)
    console.log("Posting")
    let rawResponse = await fetch(URL,
        {   
            method: 'POST',
            body: JSON.stringify(data),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            }
        }
    )

    let response = await rawResponse.json()
    console.log(response.InsertedID)
    let li = document.createElement("li")
    li.innerHTML = "<input type='checkbox' value='" + response.InsertedID +"'>"+data.itemName 
    document.querySelector("ul").appendChild(li)
    inp.value = ""
    
})



async function showList(){
    let listItems = await getItems()
    let htmlLists = listItems.map(item => {
        let li = document.createElement("li")
        li.innerHTML = "<input type='checkbox' value='" + item._id +"'>"+item.itemName
        return li
    })
    htmlLists.forEach(list => {
        document.querySelector("ul").appendChild(list)
        
    });
}
async function getItems(){
    rawResponse = await fetch(URL)
    response =  await rawResponse.json()
    return response

}

