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
  const { username, profilePicture, following, followers } = data;

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
        <div className="flex flex-col">
          <h1 className="text-3xl text-left">{username}</h1>
          <div className="flex flex-row">
            <span className="text-sm mr-3">
              <b>{followers}</b> Followers
            </span>
            <span className="text-sm">
              <b>{following}</b> Following
            </span>
          </div>
        </div>
        <div className="flex-grow" />
        <button className="btn btn-blue btn-large">
          {isLoggedIn ? "Edit Profile" : "Follow"}
        </button>
      </div>
    </div>
  );
};

export default ProfileComponent;
