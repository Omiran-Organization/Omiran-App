type Credentials = {
  username: string;
  password: string;
}

const signin = (credentials: Credentials): Promise<Response> => {
  return fetch('http://localhost:8080/signin', {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(credentials),
  })
}

const signout = (): Promise<Response> => {
  return fetch('http://localhost:8080/signout', {
    method: 'POST',
    credentials: 'include',
  })
}

export { signin, signout }
