"use client";
import React from "react";
import { WavyBackground } from "@/components/ui/wavy-background";
import Cookies from "js-cookie";

import { useState, useEffect } from "react";
import Login from "@/components/login";
import Welcome from "@/components/welcome";
import Artists from "@/components/artists-comp";

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
      <>
        {loaded ? (
          loggedIn ? (
            <div className="flex flex-wrap justify-center">
              <Welcome />
              <div className="flex flex-wrap justify-center h-full items-center">
                <div className="mx-auto">
                  <Artists />
                </div>
              </div>
            </div>
          ) : (
            <WavyBackground>
              <Login />
            </WavyBackground>
          )
        ) : (
          <div />
        )}
      </>
    </main>
  );
}