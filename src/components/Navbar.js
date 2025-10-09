import React, { useContext, useState, useRef, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import "../styles.css";
import { AuthContext } from "../context/AuthContext";

const Navbar = () => {
    const navigate = useNavigate();
    const { user, logout } = useContext(AuthContext);
    const [showUserDropdown, setShowUserDropdown] = useState(false);
    const [showUserTooltip, setShowUserTooltip] = useState(true);
    const dropdownRef = useRef(null);

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
                setShowUserDropdown(false);
                setShowUserTooltip(true);
            }
        };

        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, [dropdownRef]);

    return (
        <header className="navbar">
            <div className="logo">
                <Link to="/">DemoShop</Link>
            </div>

            <nav className="nav-links">
                <Link to="/products"><strong>Sản phẩm</strong></Link>
                <Link to="/products?category=nam"><strong>Đồ Nam</strong></Link>
                <Link to="/products?category=nu"><strong>Đồ Nữ</strong></Link>
            </nav>

            <div className="nav-icons">
                <div
                    className="icon user-menu-toggle"
                    ref={dropdownRef}
                    onClick={() => {
                        setShowUserDropdown(!showUserDropdown);
                        setShowUserTooltip(showUserDropdown); // Ẩn tooltip khi mở dropdown, hiện lại khi đóng dropdown
                    }}
                    onMouseEnter={() => {
                        if (!showUserDropdown) setShowUserTooltip(true);
                    }}
                    onMouseLeave={() => {
                        if (!showUserDropdown) setShowUserTooltip(false);
                    }}
                >
                    👤 {showUserTooltip && <span className="tooltip">{user ? "Tài khoản" : "Đăng nhập"}</span>}
                    {showUserDropdown && (
                        <div className="dropdown-menu">
                            {user ? (
                                <>
                                    <div className="dropdown-item">Chào, {user.username}!</div>
                                    {user.role === "admin" && (
                                        <div className="dropdown-item" onClick={() => {
                                            navigate("/admin");
                                            setShowUserDropdown(false);
                                            setShowUserTooltip(true);
                                        }}>Admin Dashboard</div>
                                    )}
                                    <div className="dropdown-item" onClick={() => {
                                        logout();
                                        setShowUserDropdown(false);
                                        setShowUserTooltip(true);
                                        navigate("/");
                                    }}>Đăng xuất</div>
                                </>
                            ) : (
                                <>
                                    <div className="dropdown-item" onClick={() => {
                                        navigate("/login");
                                        setShowUserDropdown(false);
                                        setShowUserTooltip(true);
                                    }}>Đăng nhập</div>
                                    <div className="dropdown-item" onClick={() => {
                                        navigate("/register");
                                        setShowUserDropdown(false);
                                        setShowUserTooltip(true);
                                    }}>Đăng kí</div>
                                </>
                            )}
                        </div>
                    )}
                </div>

                <div className="icon" onClick={() => navigate("/cart")}>
                    🛒 <span className="tooltip">Giỏ hàng</span>
                </div>
            </div>
        </header>
    );
};

export default Navbar;
