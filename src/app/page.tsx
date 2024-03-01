"use client";
import React from "react";
import { WavyBackground } from "@/components/ui/wavy-background";

import { useState, useEffect } from "react";
import Login from "@/components/login";

export default function Home() {
  const [showButton, setShowButton] = useState(false);

  return (
    <main>
      <WavyBackground className="max-w-4xl mx-auto pb-10">
        <Login />
      </WavyBackground>
    </main>
  );
}