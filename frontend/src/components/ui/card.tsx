import * as React from "react"

import { cn } from "@/lib/utils"
import Image from "next/image"

const Card = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn(
      "rounded-xl border bg-card text-card-foreground shadow",
      className
    )}
    {...props}
  />
))
Card.displayName = "Card"

const CardHeader = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("flex flex-row justify-start items-center gap-2 px-4 border-b border-gray-200", className)}
    {...props}
  />
))
CardHeader.displayName = "CardHeader"

const CardIcon = React.forwardRef<
  HTMLImageElement,
  React.ImgHTMLAttributes<HTMLImageElement>
>(({ className, alt, src, width, height, ...props }, ref) => (
  <Image
    ref={ref}
    className={cn("rounded-full", className)}
    src={src ?? ""}
    width={width as number ?? 40}
    height={height as number ?? 40}
    alt={alt ?? ""}
    {...props}
  />
))
CardIcon.displayName = "CardIcon"

const CardDescription = React.forwardRef<
  HTMLParagraphElement,
  React.HTMLAttributes<HTMLParagraphElement>
>(({ className, ...props }, ref) => (
  <p
    ref={ref}
    className={cn("font-semibold leading-none", className)}
    {...props}
  />
))
CardDescription.displayName = "CardDescription"

const CardImage = React.forwardRef <
  HTMLImageElement,
  React.ImgHTMLAttributes<HTMLImageElement>
>(({ className, alt, src, width, height, ...props }, ref) => (
  <Image
    ref={ref}
    className={cn("my-auto w-full rounded-b-xl object-cover", className)}
    src={src ?? ""}
    width={width as number ?? 500}
    height={height as number ?? 250}
    alt={alt ?? ""}
    {...props}
    />
))
CardImage.displayName = "CardImage"

const CardContent = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div ref={ref} className={cn("p-6 pt-0", className)} {...props} />
))
CardContent.displayName = "CardContent"

const CardFooter = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("flex items-center p-6 pt-0", className)}
    {...props}
  />
))
CardFooter.displayName = "CardFooter"

export { Card, CardHeader, CardFooter, CardIcon, CardDescription, CardImage, CardContent }
