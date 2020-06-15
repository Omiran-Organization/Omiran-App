import * as React from "react";

import Head from "next/head";
import { useRouter } from "next/router";
import auth from '../auth/auth-helper'
import {signin} from '../auth/api-auth'

import PasswordInput from "@/components/input/passwordinput";
import { type } from "os";

const LoginPage: React.FunctionComponent = () => {
  const router = useRouter();
  
  const [values, setValues] = React.useState({
    username: '',
    password:'',
    token: '',
    error: '',
    redirectToReferrer: false
  })
  
  // function expression not hoisted to top level scope
  const clickSubmit = () => {
    const user = {
      username: values.username || undefined,
      password: values.password || undefined
    }
    
    signin(user).then((data) => {
      console.log(`data${JSON.stringify(data)}`)
      console.log(typeof(data))
      if (data.error) {

        console.log(data.error)
      } else {
        auth.authenticate(data, () => {
          console.log(data)
          console.log(data.username,data.token)

          router.push("/profile")
         })
      }
    })
  }

  // signin(user).then((data) => {
  //   console.log(`data${JSON.stringify(data)}`)
  //   if (data.error) {

  //     setValues({ ...values,error: data.error})
  //   } else {
  //     auth.authenticate(data, () => {
  //       setValues({ ...values, error: '',redirectToReferrer: true})
  //      })
  //   }
  // })
  const handleChange = name => event => {
    
    setValues({ ...values, [name]: event.target.value })
  }

  const {redirectToReferrer} = values
    if (redirectToReferrer) {
      return router.push("/profile")
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
        {/* Email / Username
         */}
         Username
      </label>
      <input
        className="input w-full mb-3"
        type="text"
        // value={emailAddressOrUsername}
        value={values.username}
        // onChange={(e: React.ChangeEvent<HTMLInputElement>): void =>
          // setEmailAddressOrUsername(e.target.value)
        onChange={handleChange('username')}
        id="email-username-input"
      />
      <label className="w-full pl-1" htmlFor="password-input">
        Password
      </label>
      <PasswordInput
        containerProps={{ className: "mb-3" }}
        inputProps={{
          value: values.password,
          // onChange: (e: React.ChangeEvent<HTMLInputElement>): void =>
          //   setPassword(e.target.value),
          onChange: handleChange('password'),
          id: "password-input",
        }}
      />
      <button
        className="btn btn-orange self-start"
        onClick={clickSubmit}
        // onClick={(): Promise<boolean> => router.push("/profile")}
      >
        Login
      </button>
      <div className="flex-grow-3" />
    </div>
  );
};

export default LoginPage;
