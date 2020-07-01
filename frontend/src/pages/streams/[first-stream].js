// import * as React from "react";
import React, { useState, useEffect } from "react";

import Head from "next/head";

import ReactHLS from 'react-hls-player';
 

import gql from "graphql-tag";



export default function FirstPost() {

  return (
    <div className="main flex flex-col justify-center items-center w-4/5 lg:w-1/2 mx-auto text-center">
    <Head>
      <title>Omiran</title>
    </Head>
    <div className="flex mb-4">
      <div className="w-full bg-gray-500 h-12">
      <ReactHLS url={'http://localhost:8008/live/zoomer.m3u8'} controls={true} />
      </div>
    </div>
    
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