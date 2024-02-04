# Reputaton Engine

## Why Does It Exist
This is a back end microservice that provides an API to manage reputation for social media.  It's high time we had a more diverse collection of niche social media sites, to break down the monotony of monopoly, and network effect be damned.  People may want to connect but more and more they are finding that the people that connect to them are not anyone with whom they wish to associate. This is the death of big social media but still people need a way to gather on line to associate with people the DO want to associate with.

The problem is how do you moderate human behavior? Do you hire a massive number of "Moderators" do you leverage AI?  What if we stopped thinking about this as a big tech problem and started treating it as a human problem?

As annoying as I find it, features like the ones on Stack Overflow may be the answer. If you have to put in effort to build a brand on your profile, then you would be disinclined to have that profile deleted for being fake. I feel like fake profiles are currently the biggest problem on Linked In. Also if people can easily see how many people have blocked your profile as well as a pie chart showing why, so that you can tell whether they were blocked for saying things people don't like or blocked for more sinister reasons. it would be a powerful incentive to behave.

Also blocking could be reputation damage so if you lie about why you are blocking someone then you may be held accountable for libel. There needs to be a way to understand who is blocking you and why so that sort of behavior cannot happen.

## What does it do?  

This project provides a REST api which is backed by a Postgresql database. This API allows reputation management.
