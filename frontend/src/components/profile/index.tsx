import * as React from "react";

import { ProfileData } from "@/types/profile";

type ProfileComponentProps = {
  data: ProfileData;
  isLoggedIn?: boolean;
};

const ProfileComponent: React.FunctionComponent<ProfileComponentProps> = ({
  data,
  isLoggedIn,
}) => {
  const { username, profilePicture } = data;

  return (
    <div className="flex flex-col border border-gray-500 rounded-lg w-full p-5">
      <div className="flex flex-row items-center w-full">
        <img
          className="rounded-full mr-6"
          src={profilePicture}
          alt={username}
          height={100}
          width={100}
        />
        <h1 className="text-3xl">{username}</h1>
        <div className="flex-grow" />
        <button className="btn btn-blue btn-large">{isLoggedIn ? 'Edit Profile' : 'Follow'}</button>
      </div>
    </div>
  );
};

export default ProfileComponent;
