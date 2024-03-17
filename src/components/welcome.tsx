"use client";
import React from "react";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";
import ArtistCard from "./artist-card";

export default function Welcome() {
    const [username, setUsername] = useState<string | null>(null);
    const [artists, setArtists] = useState<any[] | null>(null);


    useEffect(() => {
        // use cookies to get the auth token
        fetch("http://localhost:8000/me", {
            credentials: "include"
        })
            .then((res) => {
                if (res.status === 401) {
                    window.location.href = "http://localhost:8000/login";
                }
                if (res.body === null) {
                    window.location.href = "http://localhost:8000/login";
                }
                return res.json();
            })
            .then((data) => setUsername(data.display_name));




    }, []);

    return (
        <div className="text-center">
            <h1 className="text-4xl font-bold mt-[20%]">Welcome {username}</h1>
            <p className="text-lg mb-4">Find your Favorite Artists on Vinyl</p>
        </div>
    );
}