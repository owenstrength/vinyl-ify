import React from 'react';

import Image from 'next/image';

interface ArtistCardProps {
    name: string;
    image: string;
}

const ArtistCard: React.FC<ArtistCardProps> = ({ name, image }) => {

    const handleClick = async () => {
        try {
            const response = await fetch(`http://localhost:8000/vinyl/${name}`, {
                credentials: "include"
            });
            const data = await response.json();
            // Process the data as needed
            console.log(data);
        } catch (error) {
            console.error(error);
        }
    };

    return (
        <a target="_blank" rel="noreferrer" onClick={() => handleClick()}>
            <div className="artist-card flex flex-col items-center transition-transform duration-500 transform hover:scale-125 rounded-lg">
                <Image
                    className="w-36 h-36 object-cover rounded-md"
                    src={image}
                    alt={name}
                    width={144}
                    height={144}
                />
                <h1 className="text-left font-medium font-mono">{name}</h1>
            </div>
        </a>
    );
};

export default ArtistCard;