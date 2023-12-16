"use client";

import { Button } from "@nextui-org/button";
import { Input } from "@nextui-org/input";
import { User } from "@nextui-org/react";

import { useUser } from "./hooks/useUser";

const Users = () => {
  const { username, user, isMutating, setUsername, trigger } = useUser();

  return (
    <div className="flex flex-col items-center justify-center gap-5 mt-20">
      <h1>ユーザー検索</h1>
      <div className="w-1/4">
        <Input
          value={username}
          onValueChange={setUsername}
          label="ユーザー名"
        />
      </div>
      <Button color="success" isLoading={isMutating} onClick={() => trigger()}>
        検索
      </Button>
      {user && (
        <User
          name={user.username}
          description={<p>{user.displayName}</p>}
          avatarProps={{ src: user.profileImageUrl, isDisabled: isMutating }}
        />
      )}
    </div>
  );
};

export default Users;
