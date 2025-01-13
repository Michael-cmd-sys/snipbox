var navLinks = document.querySelectorAll("nav a");
for (const i = 0; i < navLinks.length; i++) {
	const link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}