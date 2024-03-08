"use client";
import React from "react";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";

export default function Welcome() {
    const [username, setUsername] = useState<string | null>(null);
    const [artists, setArtists] = useState<Object | null>(null);


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

    useEffect(() => {
        fetch("http://localhost:8000/artists", {
            credentials: "include"
        }).then((res) => {
            if (res.status === 401) {
                window.location.href = "http://localhost:8000/login";
            }
            if (res.body === null) {
                window.location.href = "http://localhost:8000/login";
            }
            return res.json();
        })
            .then((data) => setArtists(data.items));
    }, [])

    return (
        <div className="flex justify-center items-center">
            <div className="text-center">
                <h1 className="text-4xl font-bold">Welcome {username}</h1>
                <p className="text-lg mt-4">Vinylify is a web application that allows you to create your own custom vinyl record from your favorite Spotify playlist.</p>
                <p className="text-lg mt-4">Here are some of your favorite artists: </p>
                <div className="flex justify-center items-center">
                    <div className="flex flex-wrap justify-center items-center">
                        {artists && artists.map((artist: any) => {
                            return (
                                <div key={artist.id} className="m-2 p-2 rounded-md">
                                    <p>{artist.name}</p>
                                </div>
                            )
                        })}
                    </div>
                </div>
            </div>
        </div>
    )
}