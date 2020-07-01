import React, { useState, useEffect, useRef } from 'react'

import Head from 'next/head'
import PasswordInput from '../components/passwordinput'

import { useRouter } from 'next/router'
import { signin,signout } from '../auth/api-auth'

import { UserData } from '@/types';

const LoginPage: React.FunctionComponent = () => {
  const router = useRouter()
    
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [errorMessage, setErrorMessage] = useState('')

  const hitLogin = () => {
    if (username && password) {
      tryLogin()
    } else {
      return false
    }
  }
  
  const tryLogin = () => {
    const credentials = {
      username,
      password,
    }
    
    
    signout()
    signin(credentials)
      .then(async (res) => {
        const body = await res.text()
        try {
          const data: UserData = JSON.parse(body)

          router.push(`/profile/${ data.uuid }`)
        } catch {
          setErrorMessage(body)
        }
      })
  }


  return (
    <div className="main flex flex-col items-center w-4/5 md:w-2/5 mx-auto">
      <Head>
        <title>Login - Omiran</title>
      </Head>
      <div className="flex-grow-2"/>
      <h1 className="text-4xl font-bold leading-none">Omiran</h1>
      <h3 className="text-sm mb-3">The Open Source Streaming Platform</h3>
      <div className="w-full">
        <label className="w-full pl-1" htmlFor="email-username-input">
          Username
        </label>
        <input
          className="input w-full mb-3"
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          id="email-username-input"
        />
      </div>
      <label className="w-full pl-1" htmlFor="password-input">
        Password
      </label>
      <PasswordInput
        containerProps={{ className: 'mb-3' }}
        inputProps={{
          value: password,
          onChange: (e) => setPassword(e.target.value),
          id: 'password-input',
        }}
      />
      <div className="flex flex-row items-center w-full">
        <button className="btn btn-orange self-start mr-4" onClick={tryLogin}>
          Login
        </button>
        <span className="text-red-500 capitalize">
          { errorMessage ? `Error: ${ errorMessage }` : '' }
        </span>
      </div>
      <div className="flex-grow-3" />
    </div>
  )
}

export default LoginPage