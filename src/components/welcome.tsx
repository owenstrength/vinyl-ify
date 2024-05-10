"use client";
import React from "react";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";
import ArtistCard from "./artist-card";
import Modal from "./albums-card";

export default function Welcome() {
    const [username, setUsername] = useState<string | null>(null);
    const [modalContent, setModalContent] = useState<any | null>(null);
    const [modalOpen, setModalOpen] = useState<boolean>(false);
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

    const [searchInput, setSearchInput] = useState<string>("");

    const handleSearch = async () => {
        try {
            const response = await fetch(`http://localhost:8000/vinyl?artist=${searchInput}`, {
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

    const closeModal = () => {
        setModalOpen(false);
        setModalContent(null);
    };

    return (
        <div className="text-center font-mono">
            <h1 className="text-4xl font-bold mt-[10%] ">Welcome {username}</h1>
            <p className="text-lg mb-4">Click on Your Favorite Artists</p>
            <input type="text" placeholder="Search for artists" className="rounded-lg text-black px-2 w-72" value={searchInput} onChange={(e) => setSearchInput(e.target.value)} onKeyDown={(e) => e.key === 'Enter' && handleSearch()} />

            <Modal isOpen={modalOpen} closeModal={closeModal} modalContent={modalContent} />
        </div>

    );
}