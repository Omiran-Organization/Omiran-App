import * as React from "react";

const SignupPage: React.FunctionComponent = () => (
  <div className="flex flex-col items-center h-screen w-4/5 md:w-2/5 mx-auto">
    <div className="flex-grow-2" />
    <h1 className="text-blue-500 text-4xl font-bold leading-none">Omiran</h1>
    <h3 className="text-blue-500 text-sm mb-2">
      The Open Source Streaming Platform
    </h3>
    <label className="text-blue-500 w-full pl-1" htmlFor="email-input">
      Email
    </label>
    <input className="input w-full mb-1" type="email" id="email-input" />
    <label className="text-blue-500 w-full pl-1" htmlFor="username-input">
      Username
    </label>
    <input className="input w-full mb-1" type="text" id="username-input" />
    <label className="text-blue-500 w-full pl-1" htmlFor="password-input">
      Password
    </label>
    <input className="input w-full mb-1" type="password" id="password-input" />
    <label
      className="text-blue-500 w-full pl-1"
      htmlFor="confirm-password-input"
    >
      Confirm Password
    </label>
    <input
      className="input w-full mb-2"
      type="password"
      id="confirm-password-input"
    />
    <button className="btn btn-blue self-start">Signup</button>
    <div className="flex-grow-3" />
  </div>
);

export default SignupPage;
