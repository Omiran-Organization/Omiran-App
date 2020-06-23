// import * as React from "react";
import React, { useState, useEffect } from "react";
import { initializeApollo } from "../../lib/mutations";
import Head from "next/head";



import gql from "graphql-tag";



export default function FirstPost() {

  return (
    <div className="main flex flex-col justify-center items-center w-4/5 lg:w-1/2 mx-auto text-center">
    <Head>
      <title>Omiran</title>
    </Head>
  
    <h1 className="text-orange-500 text-6xl font-bold m-3">Omiran</h1>
    <h3 className="text-white text-3xl mb-5">
      This is a standard stream
  </h3>

    <button className="bg-black px-8 py-3 text-xl rounded-lg text-white">
      Go Stream Yourself
</button>
  </div>
  );
}