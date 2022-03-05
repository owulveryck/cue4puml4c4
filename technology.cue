package c4

#Technology: {
	name: string
	type: *"" | "Db" | "Queue"
	sprite?: {
		id:  string | *"\(name)"
		url: string
	}
}

#Sprite: {
	name: string
	url:  #url
}
