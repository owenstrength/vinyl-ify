"use client";
import React from "react";
import { useEffect, useState } from "react";
import ArtistCard from "./artist-card";

export default function Artists() {
    const [artists, setArtists] = useState<any[] | null>(null);
    const [modalContent, setModalContent] = useState<any | null>(null);
    const [modalOpen, setModalOpen] = useState<boolean>(false);
    const [name, setName] = useState<string | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            if (!name) {
                return;
            }
            try {
                const response = await fetch(`http://localhost:8000/vinyl?artist=${name}`, {
                    credentials: "include"
                });
                const data = await response.json();
                if (response.ok) {
                    // Process the data as needed
                    console.log(data);
                    setModalContent(data);
                    setModalOpen(true);
                } else {
                    throw new Error('Request failed with status ' + response.status);
                }
            } catch (error) {
                console.error(error);
            }
        };

        fetchData();
    }, [name]);
    
    
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
    }, [modalContent]);

    const closeModal = () => {
        setModalOpen(false);
        setModalContent(null);
    };

    return (
        <>
        <div className="flex flex-wrap justify-center w-full">
            {artists &&
                artists.map((artist: any) => {
                    return (
                        <div key={artist.id} className="m-2 p-2">
                            <div className="rounded-md">
                                <a onClick={() => setName(artist.name)}>
                                    <ArtistCard name={artist.name} image={artist.images[0].url} />
                                </a>
                            </div>
                        </div>
                    );
                })}
        </div>
        {modalOpen && (
            <div className="fixed inset-0 flex items-center justify-center z-50">
                <div className="bg gray p-4 rounded-md">
                    <button className="absolute top-2 right-2" onClick={closeModal}>
                        Close
                    </button>
                    {modalContent && modalContent.albums.map((item: any) => {
                        return (
                            <div className="text-center font-medium font-mono text-wrap">
                                <a href={item.link} target="_blank">{item.title}</a>
                            </div>  
                        )
                    }
                )}
                </div>
            </div>
        )}
        </>
    );
}