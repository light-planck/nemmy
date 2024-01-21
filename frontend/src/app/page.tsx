"use client";

import { Button } from "@/components/ui/button";

const Home = () => {
  return (
    <div className="flex flex-col items-center justify-center gap-2 mt-9">
      <p>I&apos;m nemmy...🥱</p>
      <p className="text-2xl">Are you nemmy?</p>
      <div className="flex flex-row gap-2">
        <Button>Yes!</Button>
        <Button variant={"destructive"}>No!</Button>
      </div>
    </div>
  );
};

export default Home;
