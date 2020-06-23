// import * as React from "react";
import React, { useState, useEffect } from "react";
import { initializeApollo } from "../lib/mutations";
import Head from "next/head";
import { useRouter } from "next/router";
import auth from "../auth/auth-helper";
import { signin } from "../auth/api-auth";
import gql from "graphql-tag";
import PasswordInput from "../components/passwordinput";

import { type } from "os";

const LoginPage: React.FunctionComponent = () => {
  const router = useRouter();

  const [values, setValues] = useState({
    username: "",
    password: "",
    token: "",
    error: "",
    redirectToReferrer: false,
  });

  // function expression not hoisted to top level scope
  const clickSubmit = () => {
    const user = {
      username: values.username || undefined,
      password: values.password || undefined,
    };

    signin(user).then((data) => {
      if (data.error) {
        router.push("/profile");
      } else {
        auth.authenticate(data, () => {
          const { token } = data;

          setValues({ ...values, token: token });

          router.push("/");
        });
      }
    });
  };
  useEffect(() => {
    console.log({ ...values });
  }),
    [values];

  //getToken is (possibly) a useful helper method
  const getToken = () => {
    const token = document.cookie
      .split("; ")
      .find((row) => row.startsWith("session_token"))
      .split("=")[1];
    return token;
  };

  // handle change is a "curried function" see https://stackoverflow.com/questions/32782922/what-do-multiple-arrow-functions-mean-in-javascript
  const handleChange = (name) => (event) => {
    setValues({ ...values, [name]: event.target.value });
  };

  const { redirectToReferrer } = values;
  if (redirectToReferrer) {
    return router.push("/profile");
  }
  return (
    <div className="main flex flex-col items-center w-4/5 md:w-2/5 mx-auto">
      <Head>
        <title>Login - Omiran</title>
      </Head>
      <div className="flex-grow-2" />
      <h1 className="text-4xl font-bold leading-none">Omiran</h1>
      <h3 className="text-sm mb-3">The Open Source Streaming Platform</h3>
      <label className="w-full pl-1" htmlFor="email-username-input">
        Username
      </label>
      <input
        className="input w-full mb-3"
        type="text"
        value={values.username}
        onChange={handleChange("username")}
        id="email-username-input"
      />
      <label className="w-full pl-1" htmlFor="password-input">
        Password
      </label>
      <PasswordInput
        containerProps={{ className: "mb-3" }}
        inputProps={{
          value: values.password,
          onChange: handleChange("password"),
          id: "password-input",
        }}
      />
      <button className="btn btn-orange self-start" onClick={clickSubmit}>
        Login
      </button>
      <div className="flex-grow-3" />
    </div>
  );
};

export async function getStaticProps() {
  const apolloClient = initializeApollo();

  await apolloClient.query({


      query: gql`
        query Users {
          Users {
            uuid
            username
          }
        }
      `,
    })
    .then((result) => console.log(`result${JSON.stringify(result)}`));

  return {
    props: {
      initialApolloState: apolloClient.cache.extract(),
    },
    unstable_revalidate: 1,
  };
}

export default LoginPage;