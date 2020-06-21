export const getUsers = async () => {
    try {
      const response =  fetch("http://localhost:8080/users", {
        method: 'GET',
        credentials: 'include', 
    })
    const users = await response.json()
    return users
    } catch(err) {
      console.log(err)
    }
  }
      
