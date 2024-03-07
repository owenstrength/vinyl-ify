"use client";
import React from "react";
import { WavyBackground } from "@/components/ui/wavy-background";
import Cookies from "js-cookie";

import { useState, useEffect } from "react";
import Login from "@/components/login";

export default function Home() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [loaded, setLoaded] = useState(false);

  // look at cookies to see if user is logged in
  useEffect(() => {
    if (Cookies.get("auth_token")) {
      setLoggedIn(true);
    }
    setLoaded(true);
  }, []);


  return (
    <main>
      <WavyBackground className="max-w-4xl mx-auto pb-10">
        {loaded ? <>
          {loggedIn ? <div /> : <Login />}
        </> : <div />}
      </WavyBackground>
    </main>
  );
}