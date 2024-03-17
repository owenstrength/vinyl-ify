"use client";
import React from "react";
import { useEffect, useState } from "react";
import ArtistCard from "./artist-card";

export default function Artists() {
    const [artists, setArtists] = useState<any[] | null>(null);

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
    }, []);

    return (
        <div className="flex flex-wrap justify-center w-full">
            {artists &&
                artists.map((artist: any) => {
                    return (
                        <div key={artist.id} className="m-2 p-2">
                            <div className="rounded-md">
                                <ArtistCard name={artist.name} image={artist.images[0].url} />
                            </div>
                        </div>
                    );
                })}
        </div>
    );
}