# -*- coding: utf-8 -*-

import scrapy
import json

class ReviewSpider(scrapy.Spider):
    name = 'ReviewSpider'
    allowed_domains = ['apple.com']

    appId = ''
    authorization = ''

    maxCount = 100

    page = 0
    count = 0

    def getUrl(self, page):
        url = 'https://amp-api.apps.apple.com/v1/catalog/us/apps/{}/reviews?l=en-US&offset={}&platform=web&additionalPlatforms=appletv%2Cipad%2Ciphone%2Cmac'.format(self.appId, page*10)
        return url

    def start_requests(self):
        self.page = 0
        self.count = 0
        url = self.getUrl(self.page)
        yield scrapy.Request(
            url,
            headers={'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; Win64; x64)', 'authorization': self.authorization}
        )

    def parse(self, resp):
        results = json.loads(resp.text)
        entries = results['data']
        
        for entry in entries:
            review = {}
            attributes = entry['attributes']

            review['author'] = attributes['userName']
            review['date'] = attributes['date']
            review['rating'] = attributes['rating']
            review['title'] = attributes['title']
            review['content'] = attributes['review']

            if 'developerResponse' in attributes:
                developerResponse = attributes['developerResponse']
                if 'body' in developerResponse:
                    review['reply'] = developerResponse['body']

            yield review

        if len(entries) < 10:
            return
        
        self.count += len(entries)

        if self.count >= self.maxCount:
            return
        
        self.page += 1

        url = self.getUrl(self.page)
        yield scrapy.Request(
            url,
            headers={'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; Win64; x64)', 'authorization': 'Bearer eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IldlYlBsYXlLaWQifQ.eyJpc3MiOiJBTVBXZWJQbGF5IiwiaWF0IjoxNTkxNzI2MDEyLCJleHAiOjE2MDcyNzgwMTJ9.PF8vc_52NGR-o-E8N-kXKAuky0ikAMmBS79H0oHdbfYtXIxuqeRWhAtvNfmPTwlUs3-o2RHhxNvQGSQ46lk27w'}
        )
