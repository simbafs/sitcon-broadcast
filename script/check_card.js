// open all card and let you check they are all correct

const url = 'https://sitcon.org/2025/sessions.json'
fetch(url)
	.then(res => res.json())
	.then(d => d.sessions.map(s => s.id))
	.then(ids => ids.forEach(id => console.log(`firefox http://localhost:3000/card?id=${id}; sleep 2;`)))
