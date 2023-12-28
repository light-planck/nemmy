import { useState } from "react";
import useSWRMutation from "swr/mutation";

import { fetcher } from "@/lib/fetcher";
import { getBaseURL } from "@/lib/getBaseURL";

import { User } from "../types";

export const useUser = () => {
  const [username, setUsername] = useState("");
  const { data, isMutating, trigger } = useSWRMutation<User>(
    getBaseURL() + `/users/${username}`,
    fetcher
  );

  return {
    username,
    user: data,
    isMutating,
    setUsername,
    trigger,
  };
};
