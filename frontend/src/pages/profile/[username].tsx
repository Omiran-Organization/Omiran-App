import * as React from "react";

import Head from "next/head";
import { useRouter } from "next/router";

import ProfileComponent from "@/components/profile";

const ProfilePage: React.FunctionComponent = () => {
  const router = useRouter();
  const { username } = router.query;

  return (
    <div className="flex flex-col justify-center items-center h-screen w-4/5 lg:w-1/2 mx-auto text-center">
      <Head>
        <title>{username} - Omiran</title>
      </Head>
      <div className="flex-grow" />
      <ProfileComponent
        data={{
          username: username as string,
          profilePicture: "https://picsum.photos/100",
        }}
      />
      <div className="flex-grow-3" />
    </div>
  );
};

export default ProfilePage;
