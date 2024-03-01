"use client";
import React from "react";
import { WavyBackground } from "../components/ui/wavy-background";
import { FrontPageText } from "../components/frontpage_text";

import { useState, useEffect } from "react";
import { PinContainer } from "@/components/ui/3d-pin";

export default function Login() {
    const [showButton, setShowButton] = useState(false);

    useEffect(() => {
        const timer = setTimeout(() => {
            setShowButton(true);
        }, 4000);

        return () => clearTimeout(timer);
    }, []);

    return (
        <>
            <PinContainer>
                <a href="http://localhost:8000/login">
                    <div className="w-96 h-96" >
                        <p className="text-2xl md:text-4xl lg:text-7xl text-white font-bold inter-var text-center pt-24">
                            Vinylify
                        </p>
                        <FrontPageText />
                        <div className="flex justify-center items-center">
                            <div className={` transition-opacity ease-in-out duration-500 ${showButton ? 'opacity-100' : 'opacity-0'}`}>
                                <button
                                    className={`px-3 py-3 rounded-full bg-[#0f0f0f] font-bold text-white tracking-widest uppercase duration-500 transform transition-transform duration-500 hover:scale-105 hover:bg-[#191414] ease-in-out flex items-center text-xs sm:text-base md:text-xl `}
                                >
                                    Login with Spotify
                                    <img src="/Spotify_Icon_RGB_Green.png" alt="spotify" className="ml-2 w-6 sm:w-8" />
                                </button>
                            </div>
                        </div>
                    </div>
                </a>
            </PinContainer>

        </>
    );
}