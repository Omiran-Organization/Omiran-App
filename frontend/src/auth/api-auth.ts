type Credentials = {
  username: string;
  password: string;
}

// Removed async so that we can do error handling in component
const API_URL = process.env.API_URL ?? 'http://localhost:8080'

const signup = (credentials: Credentials & { email: string }): Promise<Response> => {
  return fetch(API_URL + '/signup', {
    method: 'POST',
    body: JSON.stringify(credentials),
  })
}

const signin = (credentials: Credentials): Promise<Response> => {
  return fetch(API_URL + '/signin', {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(credentials),
  })
}

const signout = (): Promise<Response> => {
  return fetch(API_URL + '/signout', {
    method: 'DELETE',
    credentials: 'include',
  })
}

export { signup, signin, signout }
