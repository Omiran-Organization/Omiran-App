
const signin = async (user) => {
  try {
    let response = await fetch("http://localhost:8080/signin", {
      method: 'POST',
      credentials: 'include',
      body: JSON.stringify(user)
    })
    return await response.json()
  } catch(err) {
    console.log(err)
  }
}

const signout = async () => {
  try {
    let response = await fetch('/signout/', { method: 'POST' })
    return await response.json()
  } catch(err) {
    console.log(err)
  }
}

export {
  signin,
  signout
}