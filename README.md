#dreampitcher

This code doesn't do anything important.  It's just a toy to generate some text
randomly, but it seeds from twitter searches for "i dreamed", "i wish", "i hope
i" which makes it slightly dreamy/hopeful text.  It was pretty boring until I
started trimming the first word of the text stream, which makes it slide to the
left interestingly.  Every line is tweet sized and I tried to remove all
@mentions, but very, very lazily.

This code was/is write once.  It requires some markov code ruthlessly copied
(it's BSD) from the google code walk markov example.

Respect to [Brad Garton](http://music.columbia.edu/~brad/) for introducing me
to the concept of Markov chains when I sat in on a computer music class, which
I often thought about, but didn't really think much about til the \_ebooks
craze of the last few years. 

I think this output looks oddly musical when you stream it to your console.

You'll need to move the config.gcfg.template to config.gcfg and add your
twitter API keys to make this work.  It's not truly throttled right now, so you
might hit your request limit if you just leave it running.
