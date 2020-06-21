
const signin = async (user): any => {
  try {
    const response: Response = await fetch("http://full_app:8080/signin", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify(user),
    });
    return await response.json();
  } catch(err) {
    console.log(err);
  }
};

const signout: Response = async () => {
  try {
    const response: Response = await fetch("/signout/", { method: "POST" });
    return await response.json();
  } catch(err) {
    console.log(err);
  }
};

export { signin, signout };
