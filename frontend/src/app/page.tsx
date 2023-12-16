"use client";

import { Button } from "@nextui-org/react";
import { useRouter } from "next/navigation";

const Home = () => {
  const router = useRouter();

  return (
    <div className="flex flex-col items-center justify-center gap-2 mt-9">
      <p>I&apos;m nemmy...🥱</p>
      <p className="text-2xl">Are you nemmy?</p>
      <div className="flex flex-row gap-2">
        <Button color="primary">Yes!</Button>
        <Button color="danger">No!</Button>
      </div>
      <Button color="success" onClick={() => router.push("/search/users")}>
        find users
      </Button>
    </div>
  );
};

export default Home;
