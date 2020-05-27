import * as React from "react";

import Head from "next/head";

import ProfileComponent from "@/components/profile";

const MyProfilePage: React.FunctionComponent = () => {
  const username = "Your Username Here";

  return (
    <div className="main flex flex-col justify-center items-center w-11/12 lg:w-1/2 mx-auto text-center">
      <Head>
        <title>{username} - Omiran</title>
      </Head>
      <div className="flex-grow" />
      <ProfileComponent
        data={{
          username: username as string,
          profilePicture: "https://picsum.photos/100",
          following: 5,
          followers: 10,
        }}
        isLoggedIn={true}
      />
      <div className="flex-grow-3" />
    </div>
  );
};

export default MyProfilePage;
