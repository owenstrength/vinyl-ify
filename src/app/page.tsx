"use client";
import React from "react";
import { WavyBackground } from "@/components/ui/wavy-background";
import Cookies from "js-cookie";

import { useState, useEffect } from "react";
import Login from "@/components/login";
import Welcome from "@/components/welcome";

export default function Home() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [loaded, setLoaded] = useState(false);

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
          {loggedIn ? <Welcome /> : <Login />}
        </> : <div />}
      </WavyBackground>
    </main>
  );
}