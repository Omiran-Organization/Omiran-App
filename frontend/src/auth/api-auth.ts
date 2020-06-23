type Credentials = {
  username: string;
  password: string;
}

// const signin = async (user): any => {
//   try {
//     const response: Response = await fetch("http://localhost:8080/signin", {
//       method: "POST",
//       credentials: "include",
//       body: JSON.stringify(user),
//     });
//     return await response.json();
//   } catch(err) {
//     console.log(err);
//   }
// };

// const signout: Response = async () => {
//   try {
//     const response: Response = await fetch("/signout/", { method: "POST" });
//     return await response.json();
//   } catch(err) {
//     console.log(err);
//   }
// };

// export { signin, signout };
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
