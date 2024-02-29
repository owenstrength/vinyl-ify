"use client";
import { TypewriterEffect } from "./ui/typerwriter-effect";
export function FrontPageText() {
    const words = [
        {
            text: "Find",
        },
        {
            text: "Your",
        },
        {
            text: "Favorite",
        },
        {
            text: "Music",
        },
        {
            text: "on",
        },
        {
            text: "Vinyl",
        },
    ];
    return (
        <div className="flex flex-col items-center justify-center h-[3rem]  ">
            <TypewriterEffect words={words} />
        </div>
    );
}