import React from "react";

interface ModalProps {
    isOpen: boolean;
    closeModal: () => void;
    modalContent: any;
}

const Modal: React.FC<ModalProps> = ({ isOpen, closeModal, modalContent }) => {
    if (!isOpen || !modalContent) {
        return null;
    }

    return (
        <div className="fixed inset-0 flex items-center justify-center z-50 bg-opacity-50 bg-gray-900">
            <div className="relative bg-gray-200 p-4 pt-8 rounded-md">
                <button
                    className="absolute top-1 right-2 text-gray-600 hover:text-gray-800 font-medium font-mono"
                    onClick={closeModal}
                >
                    Close
                </button>
                <h1 className="text-center font-bold font-mono text-xl text-gray-800">{modalContent.artist}</h1>
                {modalContent.albums.map((item: any, index: number) => (
                    <div
                        key={index}
                        className="text-center font-medium font-mono text-gray-800 hover:bg-gray-300 transition-colors duration-300"
                    >
                        <a href={item.link} target="_blank" rel="noopener noreferrer">
                            {item.title}
                        </a>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Modal;